import React, { useEffect, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import {
  Hide, Show, WindowFullscreen, WindowUnfullscreen,
  ScreenGetAll, WindowCenter, WindowGetSize, WindowReloadApp,
  WindowSetSize, WindowGetPosition, WindowSetPosition
} from '../../../wailsjs/runtime/runtime';

import { CaptureScreen, CaptureScreenQRCode } from '../../../wailsjs/go/crud/CrudToken';

const Capture = () => {

  const canvasRef = useRef(null);
  const divTopRef = useRef(null);
  const divLeftRef = useRef(null);
  const divRightRef = useRef(null);
  const divBottomRef = useRef(null);
  const navigate = useNavigate();

  var windowOriginSize = {w: 500, h: 350};

  var canvas    = {};
  var divTop    = {};
  var divLeft   = {};
  var divRight  = {};
  var divBottom = {};

  var ctx = {};
  var backgroundImage = {};
  var inDefineRect = false;
  var inMoveRect = false;
  var point1 = false;
  var point2 = false;

  useEffect(() => {

    canvas    = canvasRef.current;

    ctx = canvas.getContext('2d');
    ctx.imageSmoothingEnabled = false;

	  WindowGetSize().then((size) => windowOriginSize = size);

    ScreenGetAll().then(screens =>{
      for(const [index, screen] of screens.entries()){
        if(screen.isCurrent){

          canvas.width = screen.width;
          canvas.height = screen.height;

          CaptureScreen(index).then( imageData =>{

            backgroundImage = new Image();
            backgroundImage.onload = () => { 
              ctx.drawImage(backgroundImage, 0, 0, canvas.width, canvas.height);
              WindowFullscreen();
              Show();
            };
            backgroundImage.src = 'data:image/png;base64,'+imageData;
            divTop    = divTopRef.current;
            divLeft   = divLeftRef.current;
            divRight  = divRightRef.current;
            divBottom = divBottomRef.current;
          });
        }
      }
    });

    
    Hide();
  
    // capture ESC and outer keys 
    // then clear points and inDefineRect
  
    document.addEventListener('keydown', handleKeyPress);

    return () => {
      document.removeEventListener('keydown', handleKeyPress);
    }
    
  }, []);

  const handleMouseMove = (e) => {

    e.preventDefault()

    const rect = canvas.getBoundingClientRect();
    const mouseX = e.clientX - rect.left;
    const mouseY = e.clientY - rect.top;

    ctx.drawImage(backgroundImage, 0, 0, canvas.width, canvas.height);

    ctx.fillStyle = 'white';

    // Line Y
    if(!point1)ctx.fillRect( 0, mouseY, canvas.width, 0.2);

    // Line X
    if(!point2)ctx.fillRect( mouseX, 0, 0.2, canvas.height);

    if(!inDefineRect){

      ctx.font = "20px Arial";
      ctx.textAlign = "center";
      ctx.fillText("SELECT AREA",canvas.width / 2, canvas.height / 2);

    }else if( inMoveRect ){
      // re-calculate rect
      const width = Math.abs(point2.x - point1.x);
      const height = Math.abs(point2.y - point1.y);
      
      const halfWidth = width / 2;
      const halfHeight = height / 2;
    
      // Calcula os novos pontos dos vértices
      const point1X = mouseX - halfWidth;
      const point1Y = mouseY - halfHeight;
      const point2X = mouseX + halfWidth;
      const point2Y = mouseY + halfHeight;
    
      // Retorna os novos pontos em um objeto
      point1 = { x: point1X, y: point1Y };
      point2 = { x: point2X, y: point2Y };

      calcBlurRect(point1, point2);

    }else if(point1 && !point2) {

      calcBlurRect(point1, { x: mouseX , y: mouseY });

    } else if(point2){

        calcBlurRect(point1, point2);
      
    };

  };

  const calcBlurRect = (p1,p2) => {

    ctx.strokeStyle = 'white';
    ctx.setLineDash([6]);
    ctx.strokeRect( p1.x, p1.y, p2.x - p1.x, p2.y - p1.y );

//  DIV TOP and Bottom < ---      --- >
    if( p1.y < p2.y ){

      divTop.style['height']    = p1.y + 'px';

      divLeft.style['height']   = canvas.height - (p1.y + (canvas.height - p2.y) )  + 'px';
      divRight.style['height']  = canvas.height - (p1.y + (canvas.height - p2.y) )  + 'px';

      divBottom.style['height'] = (canvas.height - p2.y)  + 'px';

    } else {

      divTop.style['height'] = p2.y + 'px';

      divLeft.style['height']   = canvas.height - (p2.y + (canvas.height - p1.y) )  + 'px';
      divRight.style['height']  = canvas.height - (p2.y + (canvas.height - p1.y) )  + 'px';
      
      divBottom.style['height'] = (canvas.height - p1.y)  + 'px';

    }

//  DIV < --- >  < --- >
    if( p1.x < p2.x ){
      divLeft.style['width']   = p1.x + 'px';
      divRight.style['width']  = canvas.width - p2.x + 'px';
    } else {
      divLeft.style['width']   = p2.x + 'px';
      divRight.style['width']  = canvas.width - p1.x + 'px';
    }

  }

  const handleMouseClickDefineRect = (e) => {

    e.preventDefault()
    e.stopPropagation()

    const rect = canvas.getBoundingClientRect();
    const mouseX = e.clientX - rect.left;
    const mouseY = e.clientY - rect.top;
    
    inDefineRect = true;

    if(!point1){
      point1 = { x: mouseX, y: mouseY };
    } else if(!point2){
      point2 = { x: mouseX, y: mouseY };
    // reset
    } else {
      
      point1 = { x: mouseX, y: mouseY };
      point2 = false;
      calcBlurRect(point1, point1);
    }

    
    if(point1 && point2) CaptureScreenQRCode([Math.ceil(point1.x), Math.ceil(point1.y)],[Math.ceil(point2.x), Math.ceil(point2.y)])
    .then(handlerCaptureScreenQRCode);


  }; 

  // func move rect for position clicked and mause position move
  const handleClickInRectMove = (e)=>{

    if( !point1 || !point2 )return;
    

    e.preventDefault();
    e.stopPropagation();
    
    inMoveRect = !inMoveRect;


    if(!inMoveRect) CaptureScreenQRCode([Math.ceil(point1.x), Math.ceil(point1.y)],[Math.ceil(point2.x), Math.ceil(point2.y)])
    .then(handlerCaptureScreenQRCode);

  }
  const handlerCaptureScreenQRCode = (result_raw) =>{
    console.log(result_raw)
    let result = JSON.parse(result_raw)

    if(result.status){

      navigate("/token/create/", { state: result.message })

      handleWindowReset().then(() => {}).catch(()=> {})

    // no detect
    } else {

      ctx.strokeStyle = 'red';
      ctx.setLineDash([6]);
      ctx.strokeRect( point1.x, point1.y, point2.x - point1.x, point2.y - point1.y );  

    }
  }

  const handleKeyPress = (event) => {

    if (
        event.keyCode == 27 || // ESC
        event.keyCode == 32 || // ESPACE
        event.keyCode == 17 || // CTRL
        event.keyCode == 16  // SHIFT
      ) {

      	// se ao tentar pegar o position eu ja tenha maximizado em fullcreen a tela?
		  setTimeout(() => handleWindowReset().then(() => {}).catch(()=> WindowReloadApp()), 50)

      navigate("/")
	}

  };

  const handleWindowReset = () => {

	  WindowUnfullscreen();
    
	  return new Promise((resolve,reject)=>{
      WindowCenter()
      WindowSetSize(windowOriginSize.w , windowOriginSize.h).then(size=>resolve())
      .catch(e=>reject(e)); 
    });

  }

  const styles = {
    container:{
      position:'absolute',
      width:'100vw',
      height:'100vh',
	  overflow: 'hidden',
      top:'0',
      left:'0'
    },
    canvas: {
      backgroundColor: 'lightblue',
      margin:'0px',
      padding:'0px',
      display: 'block',
      width:'100vw',
      height:'100vh',
      position:'absolute',
      top:'0',
      left:'0'
    },
    divTop:{
      'background': '#101010bf',
      'height':'100vh',
      'width':'100%',
      'float':'left'
    },
    span:{
      'width':'100vw',
      'display':'block',
      'overflow':'auto'
    },
    divRight:{
      'background': '#101010bf',
      'width':'50%',
      'float':'right'
    },
    divLeft:{
      'background': '#101010bf',
      'width':'50%',
      'float':'left'
    },
    divBottom:{
      'background': '#101010bf',
      'width':'100%',
      'float':'left'
    }
  };

  return (
    <div>
      <canvas style={styles.canvas}     ref={canvasRef} ></canvas>
      <div onClick={handleMouseClickDefineRect} onMouseMove={handleMouseMove} style={styles.container}>
        <div  style={styles.divTop} ref={divTopRef}></div>
        <span onClick={handleClickInRectMove} style={styles.span}>
          <div onClick={handleMouseClickDefineRect} style={styles.divRight}  ref={divRightRef}></div>
          <div onClick={handleMouseClickDefineRect} style={styles.divLeft}   ref={divLeftRef}></div>
        </span>
        <div  style={styles.divBottom}   ref={divBottomRef}></div>
      </div>
  \  </div>
  );
};

export default Capture;