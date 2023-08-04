import React, { useEffect, useRef } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';


const Password = () => {

  const navigate = useNavigate();
  const location = useLocation();
  const inputPass = useRef(null);

  const handlerSubmit = (event) =>{
    
    event.preventDefault()

    let data = location.state
    data.pass = inputPass.current.value +""
    console.log(data);
    navigate( data.dst, { state: data });

  }

  return (
    <div>
      <h1>Token Password</h1>
      <form onSubmit={handlerSubmit}>
        <label htmlFor="pass">Password</label>
        <input ref={inputPass} type="password" name="pass"/>
        <input type="submit" value="Unlock"/>
      </form>
    </div>
  );
};

export default Password;