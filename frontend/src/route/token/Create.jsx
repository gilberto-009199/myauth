import React,{ useEffect, useState } from 'react';

const Create = (props) => {
  
  

  const [OTP, setOTP] = useState("");

  useEffect(() => {

    console.log("props",props)
    console.log("props",props.state)
    

    //setOTP(JSON.parse(props.match.params.otp));
      
  }, []);

  return (
    <div>
      <h1>Token Create</h1>
    </div>
  );
};

export default Create;