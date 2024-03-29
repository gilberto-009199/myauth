import React, { useEffect, useRef, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

import { ImportToken } from '../../../wailsjs/go/main/App';

const Import = () => {

  const navigate = useNavigate()
  const location = useLocation()

  const [token, setToken] = useState({})

  const formatRef = useRef(null);

  useEffect(() => {

    if( location.state ){
      console.log(location.state)
      if(typeof location.state.pass != 'undefined'){

        formatRef.current.value = location.state.format
        
        handlerExport()

      } else {
        location.state.dst = '/token/import'
        navigate('/token/password', { state: location.state });
      }
    }

  },[])

  const handlerExport = (event = false) =>{
    
    if(event)event.preventDefault()

    console.log("importacao!!")
    let sltValue = formatRef.current.value

    if( sltValue != "myauth" || (location.state && typeof location.state.pass != 'undefined') ){

      console.log("Iniciar importacao!!")
      ImportToken(sltValue, token.pass).then(res=>{
        let data = JSON.parse(res)
        if(data.status){
          navigate('/');
        }else console.log("ERRO", res)
      }).catch(e=>console.log(e))

    } else {
      navigate('/token/password', { state: {
        dst:'/token/import',
        format: sltValue
      } });
    }

    
  }

  return (
    <div className='router-content'>
      <form onSubmit={handlerExport}>
        <h1>Token Import</h1>
        <select ref={formatRef}>
          <option value="myauth">MyAuth</option>
          <option value="csv">CSV</option>
          <option value="qrcode">QRCode</option>
        </select>
        <button type='submit'>import</button>
      </form>
    </div>
  );
};

export default Import;