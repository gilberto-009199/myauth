import React,{ useEffect, useState, useRef } from 'react';
import { useLocation, useNavigate, Link } from 'react-router-dom';

import { TokenList } from '../../wailsjs/go/crud/CrudToken';


import '../assets/css/home.css';

const Home = () => {

  const location = useLocation();
  const navigate = useNavigate();
  const divListTokens = useRef(null);
  
  const [listTokens, setListTokens] = useState([]);
  
  var passwrd = { pass:"", time: 0 }

  useEffect(() => {

    loadTokens()

    let idTimer = setInterval(()=>{ loopTokens() }, 150)

    return ()=>{
      clearInterval(idTimer)
    }

  }, []);

// ((100 / 30).toFixed(1) * 30 -1) = circle.style.strokeDashoffset
// document.querySelector('.circle.circle_animation')
// circle.style.strokeDashoffset = '50'
// Math.floor( Date.now() / (30 * 1000) ) <== secound loop
// ( Date.now() / (30 * 1000)) - Math.floor( Date.now() / (30 * 1000) )

  const loopTokens = () =>{

    let timeNow = Date.now()

    divListTokens.current.querySelectorAll('.circle.circle_animation').forEach( elm =>{

      let interval = elm.attributes['data-second'].value * 1;
      let text  = elm.parentNode.querySelector('text')

      let percentage = ( ( timeNow / (interval * 1000)) - Math.floor( timeNow / (interval * 1000) ) )* 100

      elm.style.strokeDashoffset = percentage 

      if(percentage > 98){

        text.innerHTML = (0) + 's';

        setTimeout(()=> loadTokens(), 100)

      } else text.innerHTML = ~(( interval / 100 * percentage ).toFixed(0)  - interval) + 's';

    })
  }

  const loadTokens = () =>{
    TokenList(passwrd.pass).then(res =>{
      let data = JSON.parse(res)
      if(data.status){
        setListTokens(data.message)
      }
    });
  }

  const maskCode = (code) =>{
    return code.slice(0, code.length / 2) + " " + code.slice( code.length / 2,  code.length )
  }

  const toogleCRUD = (event) =>{
    
    let itemDOM = event.target;

    if(itemDOM.parentNode.parentNode.parentNode.querySelector(".list_token_item_menu").classList.toggle("disable")){
      itemDOM.style.transform = 'rotate(0deg)'
    }else{
      itemDOM.style.transform = 'rotate(90deg)'
    }

  }
  return (
    <div>
      <div className='router-content'>
      {listTokens.length === 0 ? (
        <div>
          <h2 data-icon="&#x71;"></h2>
          <p>Welcome MyAuth</p>
        </div>
      ) : (
        <div className="list_token" ref={divListTokens}>
          {listTokens.map((item, index) => (
            <div key={index} className='list_token_item' >
              <svg className="list_token_item_timer" width="50" height="50" xmlns="http://www.w3.org/2000/svg">
                <g>
                  <circle r="16" cy="25" cx="25" strokeWidth="2" stroke="#003fff75" fill="none"></circle>
                  <circle className="circle circle_animation" data-second="30" r="16" cy="25" cx="25" strokeWidth="3" stroke="#6fdb6f" fill="none"></circle>
                  <text className="list_token_item_timer_text" x="25" y="30">30s</text>
                </g>
              </svg>
              <div>
                <p className='list_token_item_name'>{item.name}</p>
                <p className='list_token_item_code'>
                  { maskCode(item.code) }
                  <span className="list_token_item_btn" data-icon="&#x35;" onClick={toogleCRUD}></span>
                </p>
              </div>
              <div className='list_token_item_menu disable'>
                <div>
                  <Link className='borda' to={"/token/update"} data-icon="&#xe035;" state={item}>Update</Link>
                  <Link className='borda' to={"/token/info"} data-icon="&#xe060;" state={item}>Info</Link>
                  <Link className='borda' to={"/token/remove"} data-icon="&#xe051;" state={item}>Remover</Link>
                  <Link to={"/token/export"} data-icon="&#xe05e;" state={item}>Exportar</Link>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
      </div>
      <div className='panel-action-buttons'>
          <button className='dark' style={{'width':'48%'}} onClick={()=> navigate("/token/import")}>
            Use existing Token
          </button>
          <button style={{'width':'48%'}} onClick={()=> navigate("/token/create")}>
            Create new Token
          </button>
      </div>
    </div>
  );
};

export default Home;