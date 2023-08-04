import React,{ useEffect, useState } from 'react';
import { useLocation, useNavigate, Link } from 'react-router-dom';

import { TokenList } from '../../wailsjs/go/crud/CrudToken';

const Home = () => {

  const location = useLocation();
  const navigate = useNavigate();
  
  const [listTokens, setListTokens] = useState([]);
  
  var passwrd = { pass:"", time: 0 }

  useEffect(() => {
    
    TokenList(passwrd.pass).then(res =>{
      res = JSON.parse(res)
      if(res.status){
        setListTokens(res.message)
      }
    });

  }, []);

  return (
    <div>
      <h1>Home Page</h1>
      <div>
      {listTokens.length === 0 ? (
        <div>
          <Link to={"/token/capture"} data-icon="&#xe050;"></Link>
          <p>Adicione itens à lista</p>
        </div>
      ) : (
        <div>
          {listTokens.map((item, index) => (
            <div key={index}>
              <label>{item.name} - {item.code}</label>
              <div>
                <Link to={"/token/update"} state={item}>Edit</Link><br/>
                <Link to={"/token/info"} state={item}>Info</Link><br/>
                <Link to={"/token/remove"} state={item}>remove</Link><br/>
                <Link to={"/token/export"} state={item}>export</Link><br/>
              </div>
            </div>
          ))}
        </div>
      )}
      </div>
    </div>
  );
};

export default Home;