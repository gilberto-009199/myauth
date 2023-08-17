import React,{ useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useFormik } from 'formik';
import * as Yup from 'yup';

import { ListAlgoritm  as getListAlgoritm , SetSettings } from '../../wailsjs/go/main/App';
import { GetSettings, GetFile } from '../../wailsjs/go/main/App';

const Settings = () => {

  const navigate = useNavigate()

  const [listAlgoritm, setListAlgoritm] = useState([]);

  getListAlgoritm().then(res =>{
    let data = JSON.parse(res)
    setListAlgoritm( data.message )
  });

  useEffect(()=>{ 

    GetSettings().then(res=>{
      let data = JSON.parse(res);

      formik.setValues({
        file_settings: data.message.file_settings,
        file_tokens: data.message.file_tokens,
        algoritm: data.message.algoritm
      })

    })

   },[]);

  const formik = useFormik({
    initialValues: {
      file_settings: '',
      file_tokens: '',
      algoritm:''
    },
    validationSchema: Yup.object({
      file_settings: Yup.string().required('The file is necessary').min(3),
      file_tokens: Yup.string().required('The file is necessary').min(3),
      algoritm: Yup.string().required('The algoritm is necessary')
    }),
    onSubmit: (values) => {
      // Realizar ações de envio do formulário
      console.log('Dados do formulário:', values);
      SetSettings(JSON.stringify(values)).then(res=>{
        console.log(res)
        let data = JSON.parse(res)
        if(data.status){
          navigate('/')
        }
      }).catch(e => console.log(e))
      // Limpar os campos após o envio
      formik.resetForm();
    },
  });

  const clickDiscSettiing = () => {
    GetFile('*.*').then(res => {
      let data = JSON.parse(res)
      if(data.status){
        formik.setValues({
          file_settings: data.message,
          file_tokens: formik.values.file_tokens,
          algoritm: formik.values.algoritm
        })
      }
    }).catch(e=>console.log(e))
  }

  const clickDiscToken = () => {
    GetFile('*.bin;*.*').then(res => {
      let data = JSON.parse(res)
      if(data.status){

        formik.setValues({
          file_settings: formik.values.file_settings,
          file_tokens: data.message,
          algoritm: formik.values.algoritm
        })

      }
    }).catch(e=>console.log(e))
  }

  return (
    <div className='router-content'>
      <form onSubmit={formik.handleSubmit}>
        <h1>Settings</h1>
        <div>
          <label htmlFor="file_settings">File Settings:</label>
          <br />
          <input
            type="text"
            name="file_settings"
            value={formik.values.file_settings}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
          />
          <button type='button' className='input_complement' data-icon="&#xe0e4;" onClick={clickDiscSettiing}></button>
          {formik.touched.file_settings && formik.errors.file_settings && (
            <span className='error-input'>{formik.errors.file_settings}</span>
          )}
        </div>
        <div>
          <label htmlFor="file_tokens">File Settings:</label>
          <br />
          <input
            type="text"
            name="file_tokens"
            value={formik.values.file_tokens}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}/>
          <button type='button' className='input_complement' data-icon="&#xe0e4;" onClick={clickDiscToken}></button>
          {formik.touched.file_tokens && formik.errors.file_tokens && (
            <span className='error-input'>{formik.errors.file_tokens}</span>
          )}
        </div>
        <div>
          <label htmlFor="algoritm">Algoritm:</label>
          <br />
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
        <button type="submit">Salvar</button>
      </form>
    </div>
  );
};

export default Settings;