import React, { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { TokenDelete } from '../../../wailsjs/go/crud/CrudToken';

const Remove = () => {

  const [token, setToken] = useState({
    name:'NameToken',
  })
  const location = useLocation()
  const navigate = useNavigate()

  useEffect(()=>{
    if(location.state){

      setToken(location.state)

    }else navigate(-1)

  },[])

  const handlerRemoveToken = () => {
    TokenDelete(token.id).then(res=>{
      console.log(res)
      let data = JSON.parse(res);
      if(data.status){
        navigate('/')
      } else {
        // menssage error
      }

    })
  }

  return (
    <div>
      <h1>Token Remove</h1>
      <p>Deseja mesmo remover o Token? {token.name}</p>
      <button onClick={ handlerRemoveToken }>Sim</button>
      <button onClick={ () => navigate(-1) }>Não</button>
    </div>
  );
};

export default Remove;