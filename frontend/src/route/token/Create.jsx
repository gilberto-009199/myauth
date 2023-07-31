import React,{ useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { useFormik } from 'formik';
import * as Yup from 'yup';
import { ListAlgoritm  as getListAlgoritm } from '../../../wailsjs/go/main/App';
import { TokenCreate } from '../../../wailsjs/go/crud/CrudToken';

const Create = (props) => {

  // feature add campo code current for register
  const location = useLocation();
  
  const [listAlgoritm, setListAlgoritm] = useState([]);
  
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

  }, []);

  return (
    <div>
      <h1>Token Create</h1>
      <form onSubmit={formik.handleSubmit}>
        <div>
          <label>Name:</label>
          <input
            type="text"
            name="name"
            value={formik.values.name}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
          />
          {formik.touched.name && formik.errors.name && (
            <span>{formik.errors.name}</span>
          )}
        </div>
        <div>
          <label>Secret:</label>
          <input
            type="password"
            name="secret"
            value={formik.values.secret}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
          />
          {formik.touched.secret && formik.errors.secret && (
            <span>{formik.errors.secret}</span>
          )}
        </div>
        <div>
          <label>Algoritm:</label>
          <select
            name="algoritm"
            value={formik.values.algoritm}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            >
              <option value="">Selecione um item</option>
              {
                listAlgoritm.map((item,index) => (
                  <option key={index} value={item}>{item}</option>
                ))
              }
          </select>
          {formik.touched.algoritm && formik.errors.algoritm && (
            <span>{formik.errors.algoritm}</span>
          )}
        </div>
        <button type="submit">Salvar</button>
      </form>
    </div>
  );
};

export default Create;