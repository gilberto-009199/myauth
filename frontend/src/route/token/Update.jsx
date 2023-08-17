import React,{ useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { useFormik } from 'formik';
import * as Yup from 'yup';

import { ListAlgoritm  as getListAlgoritm } from '../../../wailsjs/go/main/App';
import { TokenUpdate } from '../../../wailsjs/go/crud/CrudToken';

const Update = () => {

  const navigate = useNavigate()
  const location = useLocation()
  const [listAlgoritm, setListAlgoritm] = useState([]);
   
  getListAlgoritm().then(res =>{
    let data = JSON.parse(res)
    setListAlgoritm( data.message )
  });

  useEffect(()=>{

    if( location.state ){
      console.log("DADOS DO UPDATE",location.state)
      
      if(typeof location.state.pass != 'undefined'){

        formik.setValues({
          name: location.state.name,
          algoritm: location.state.algoritm,
          url: location.state.url
        })

      } else {
        location.state.dst = '/token/update'
        navigate('/token/password', { state: location.state });
      }
      

    } else navigate('/', { state: location.state })

  },[])

  const formik = useFormik({
    initialValues: {
      name: '',
      algoritm: '',
      url:''
    },
    validationSchema: Yup.object({
      name: Yup.string().required('The name is necessary').min(3),
      algoritm: Yup.string().required('The algoritm is necessary')
    }),
    onSubmit: (values) => {
      // Realizar ações de envio do formulário
      console.log('Dados do formulário:', values);
      
      TokenUpdate(location.state.id, JSON.stringify(values), location.state.pass).then(res=>{
        console.log(res)
        navigate('/')
      }).catch(e=>{
        console.log(e)
      })
      // Limpar os campos após o envio
      formik.resetForm();
    },
  });

  return (
    <div className='router-content'>
      <form onSubmit={formik.handleSubmit}>
        <div>
          <label htmlFor="name">Name:</label>
          <br />
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
        <div>
          <label htmlFor="algoritm">Algoritm:</label>
          <br />
          <select
            name="algoritm"
            value={formik.values.algoritm}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            >
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
        <button type="submit">Salvar</button>
      </form>
    </div>
  );
};

export default Update;