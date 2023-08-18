import React,{ useEffect, useState } from 'react';
import { useLocation, useNavigate, Link } from 'react-router-dom';

import { useFormik } from 'formik';
import * as Yup from 'yup';
import { ListAlgoritm  as getListAlgoritm } from '../../../wailsjs/go/main/App';
import { TokenCreate, TokenTimeCode } from '../../../wailsjs/go/crud/CrudToken';


import '../../assets/css/token.css';

const Create = (props) => {

  // feature add campo code current for register
  const location = useLocation();
  const navigate = useNavigate();
  const [listAlgoritm, setListAlgoritm] = useState([]);
  const [genCodeValid, setGenCodeValid] = useState(false);
  const [currentToken, setCurrentToken] = useState('000000');
  const [nextToken, setNextToken] = useState('000000');
  
  getListAlgoritm().then(res =>{
    let data = JSON.parse(res)
    setListAlgoritm( data.message )
  });

  const formik = useFormik({
    initialValues: {
      name: '',
      secret: '',
      algoritm: '',
      url:''
    },
    validationSchema: Yup.object({
      name: Yup.string().required('The name is necessary').min(3),
      secret: Yup.string().required('The secret is necessary').min(4),
      algoritm: Yup.string().required('The algoritm is necessary')
    }),
    onSubmit: (values) => {
      // Realizar ações de envio do formulário
      console.log('Dados do formulário:', values);
      TokenCreate(JSON.stringify(values)).then(res=>{
        console.log(res)
        navigate('/')
      }).catch(e=>{
        console.log(e)
      })
      // Limpar os campos após o envio
      formik.resetForm();
    },
  });

  useEffect(() => {

    if( location.state ){
        console.log(location.state)
        
        formik.setValues({
          name: location.state.name || '',
          secret: location.state.secret || '',
          algoritm: location.state.algoritm,
          url: location.state.url
        })

    }

    let idTimer = setInterval(()=>{ loopTokens() }, 150)

    return ()=>{
      clearInterval(idTimer)
    }

  }, []);

  const loopTokens = () =>{

    let timeNow = Date.now()

    document.querySelectorAll('.circle.circle_animation').forEach( elm =>{

      let interval = elm.attributes['data-second'].value * 1;
      let text  = elm.parentNode.querySelector('text')

      let percentage = ( ( timeNow / (interval * 1000)) - Math.floor( timeNow / (interval * 1000) ) )* 100

      elm.style.strokeDashoffset = percentage * 2

      if(percentage > 98){

        text.innerHTML = (0) + 's';

        setTimeout(()=> handlerToken({target:{ value: formik.values.secret }}), 100)

      } else text.innerHTML = ~(( interval / 100 * percentage ).toFixed(0)  - interval) + 's';

    })
  }

  const maskCode = (code) =>{
    return code.slice(0, code.length / 2) + " " + code.slice( code.length / 2,  code.length )
  }
  const handlerToken = (e)=>{
    
    let code = e.target.value

    if( code.length < 4 )return;

    TokenTimeCode(code).then(res=>{
      let data = JSON.parse(res)
      if(data.status){
        setCurrentToken(data.message[0])
        setNextToken(data.message[1])
        setGenCodeValid(true)
      }else setGenCodeValid(false)
    })

  }


  return (
    <div className='router-content'>
      <div className="painel-create-captureqr" onClick={()=> navigate("/token/capture",{ state: {'dst':'/token/create'} }) }>
        <Link data-icon="&#xe00f;"></Link>
        <br/>
        QRCode
      </div>
      
      <form className="form-create" onSubmit={formik.handleSubmit}>
        <div className="form_group">
          <label htmlFor="name">Name:</label>
          <input
            type="text"
            name="name"
            value={formik.values.name}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
          />
          {formik.touched.name && formik.errors.name && (
            <span className='error-input'>{formik.errors.name}</span>
          )}
        </div>
        <div className="form_group">
          <label htmlFor="secret">Secret:</label>
          <input
            type="password"
            name="secret"
            onInput={handlerToken}
            value={formik.values.secret}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
          />
          {formik.touched.secret && formik.errors.secret && (
            <span className='error-input'>{formik.errors.secret}</span>
          )}
        </div>
        <div className="form_group">
          <label htmlFor="algoritm">Algoritm:</label>
          <select
            name="algoritm"
            value={formik.values.algoritm}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}>
              {
                listAlgoritm.map((item,index) => (
                  <option key={index} value={item}>{item}</option>
                ))
              }
          </select>
          {formik.touched.algoritm && formik.errors.algoritm && (
            <span className='error-input'>{formik.errors.algoritm}</span>
          )}
        </div>
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
        <button type="submit">Salvar</button>
      </form>
    </div>
  );
};

export default Create;