import React,{ useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

import '../../assets/css/token.css';

const Info = () => {

  const location = useLocation()
  const navigate = useNavigate()

  const [genCodeValid, setGenCodeValid] = useState(false);
  const [currentToken, setCurrentToken] = useState('000000');
  const [nextToken, setNextToken] = useState('000000');

  const [token, setToken] = useState({
    name: '',
    secret: '***',
    algoritm: '',
    code: '***'
  })

  useEffect(()=>{

    if( location.state ){
 
      if(typeof location.state.pass != 'undefined'){
        // solicitar decript 
        var idTimer = setInterval(()=>{ loopTokens() }, 150)

      }

      setToken({
        name: location.state.name,
        secret: location.state.secret,
        algoritm: location.state.algoritm,
        code: location.state.code
      })

    } else navigate('/', { state: location.state })
    return ()=>{
      if(typeof location.state.pass != 'undefined')clearInterval(idTimer)
    }
  },[])

  const loopTokens = () =>{

    let timeNow = Date.now()

    document.querySelectorAll('.circle.circle_animation').forEach( elm =>{

      let interval = elm.attributes['data-second'].value * 1;
      let text  = elm.parentNode.querySelector('text')

      let percentage = ( ( timeNow / (interval * 1000)) - Math.floor( timeNow / (interval * 1000) ) )* 100

      elm.style.strokeDashoffset = percentage * 2

      if(percentage > 98){

        text.innerHTML = (0) + 's';

        setTimeout(()=> handlerToken(), 100)

      } else text.innerHTML = ~(( interval / 100 * percentage ).toFixed(0)  - interval) + 's';

    })
  }
  const maskCode = (code) =>{
    return code.slice(0, code.length / 2) + " " + code.slice( code.length / 2,  code.length )
  }

  const handlerToken = ()=>{
    
    console.log(location.state)

    let Id = location.state.id

    /*
    
    AQUIIIII
    TokenTimeIdToken(id).then(res=>{
      let data = JSON.parse(res)
      if(data.status){
        setCurrentToken(data.message[0])
        setNextToken(data.message[1])
        setGenCodeValid(true)
      }else setGenCodeValid(false)
    })*/

  }

  return (
    <div className='router-content'>
      <div>
        <div className="token_desc">
          { genCodeValid ?
          <div className="token_code_timers" >
            <div className='token_code_timer'>
              <svg className="list_token_item_timer" width="100" height="100" xmlns="http://www.w3.org/2000/svg">
                  <g>
                    <circle r="32" cy="50" cx="50" strokeWidth="2" stroke="#003fff75" fill="none"></circle>
                    <circle className="circle circle_animation" data-second="30" r="32" cy="50" cx="50" strokeWidth="3" stroke="#6fdb6f" fill="none"></circle>
                    <text className="list_token_item_timer_text" x="50" y="54">30s</text>
                  </g>
              </svg>
              {maskCode(currentToken)}
            </div>
            <div className='token_code_timer'>
            <svg className="list_token_item_timer" width="100" height="100" xmlns="http://www.w3.org/2000/svg">
                  <g>
                    <circle r="32" cy="50" cx="50" strokeWidth="2" stroke="#003fff75" fill="none"></circle>
                    <circle className="circle circle_animation" data-second="30" r="32" cy="50" cx="50" strokeWidth="3" stroke="#6fdb6f" fill="none"></circle>
                    <text className="list_token_item_timer_text" x="50" y="54">30s</text>
                  </g>
              </svg>
              {maskCode(nextToken)}
            </div>
          </div>
          : 
          <div className="token_code_noffound">
            <p data-icon="&#x71;"></p>  
            <p> Codigo não reconhecido! </p>
          </div>}
        </div>
        <div className="token_desc">
          Name: 
          <div className="token_desc_value">{token.name}</div>
        </div>
        <div className="token_desc">
          Secret:
          <div className="token_desc_value">{token.secret} Ver </div>
        </div>
        <div className="token_desc">
          Algoritm:
          <div className="token_desc_value">{token.algoritm}</div>
        </div>
      </div>
    </div>
  );
};

export default Info;