import React,{ useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import { useFormik } from 'formik';
import * as Yup from 'yup';
import { ListAlgoritm  as getListAlgoritm } from '../../../wailsjs/go/main/App';

const Create = () => {

  // feature add campo code current for register
  const location = useLocation();
  
  const [listAlgoritm, setListAlgoritm] = useState([]);
  
  getListAlgoritm().then(res =>{
    let data = JSON.parse(res)
    setListAlgoritm(data['Message'])
  });

  const formik = useFormik({
    initialValues: {
      name: '',
      secret: '',
      algoritm: '',
    },
    validationSchema: Yup.object({
      name: Yup.string().required('The name is necessary').min(3),
      secret: Yup.string().required('The secret is necessary').min(4),
      algoritm: Yup.string().required('The algoritm is necessary')
    }),
    onSubmit: (values) => {
      // Realizar ações de envio do formulário
      console.log('Dados do formulário:', values);

      // Limpar os campos após o envio
      formik.resetForm();
    },
  });

  useEffect(() => {

    if( location.state ){
      
        formik.values.name = location.state.name || '';
        formik.values.secret = location.state.secret || '';
        formik.values.algoritm = location.state.algoritm || '';
        formik.values.url = location.state.url || '';

    }
    console.log(listAlgoritm);

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
            onBlur={formik.handleBlur} >
              {
                listAlgoritm.map((item, index) => (
                  <option value={index}>{item}</option>
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