import React,{ useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';


const Info = () => {

  const location = useLocation()
  const navigate = useNavigate()

  const [token, setToken] = useState({
    name: '',
    secret: '***',
    algoritm: '',
    code: '***'
  })

  useEffect(()=>{

    if( location.state ){
      console.log("DADOS DO INFO",location.state)
      
      if(typeof location.state.pass != 'undefined'){
        // solicitar decript 
      }

      setToken({
        name: location.state.name,
        secret: location.state.secret,
        algoritm: location.state.algoritm,
        code: location.state.code
      })

    } else navigate('/', { state: location.state })

  },[])

  return (
    <div>
      <h1>Token Info</h1>
      <p>Name: <strong>{token.name}</strong></p>
      <p>Secret: <strong>{token.secret}</strong> Ver </p>
      <p>Algoritm: <strong>{token.algoritm}</strong></p>
      <p>Code: <strong>{token.code}</strong></p>
    </div>
  );
};

export default Info;