import React, { useEffect, useRef, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

import { ExportToken } from '../../../wailsjs/go/main/App';

const Export = () => {

  const navigate = useNavigate()
  const location = useLocation()

  const [token, setToken] = useState({})

  const formatRef = useRef(null);

  useEffect(() => {

    if( location.state ){

      console.log("DADOS DO Export",location.state)
      
      if(typeof location.state.pass != 'undefined'){

        setToken(location.state)

      } else {
        location.state.dst = '/token/export'
        navigate('/token/password', { state: location.state });
      }

    } else navigate('/', { state: location.state })

  },[])

  const handlerExport = (event) =>{
    event.preventDefault()

    let sltValue = formatRef.current.value

    ExportToken(token.id, sltValue, token.pass).then(res=>{
      let data = JSON.parse(res)
      
      if(data.status){
        navigate('/');
      }else console.log("ERRO", res)

    }).catch(e=>console.log(e))

  }

  return (
    <div  className='router-content'>
      <form onSubmit={handlerExport}>
        <h1>Token {token.name} Export</h1>
        <select ref={formatRef}>
          <option value="myauth">MyAuth</option>
          <option value="csv">CSV</option>
          <option value="qrcode">QRCode</option>
        </select>
        <button>Export</button>
      </form>
    </div>
  );
};

export default Export;