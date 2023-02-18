import React, { useEffect, useState } from 'react'
import qs from 'qs'
import './Home.css'
import axios from 'axios'

function Login() {
  const [state, setState] = useState('')
  let loginUrl = ''

  useEffect(() => {
    axios
      .get(
        `${process.env.REACT_APP_SERVER_ORIGIN as string}/oauth/google/state`,
      )
      .then((Response) => {
        setState(Response.data.state)
      })
      .catch((error) => {
        console.log(error)
      })
  }, [])

  if (state !== '') {
    const AUTHORIZE_URI = 'https://accounts.google.com/o/oauth2/v2/auth'
    const queryStr = qs.stringify({
      client_id: process.env.REACT_APP_OAUTH_GOOGLE_CLIENT_ID as string,
      redirect_uri: process.env.REACT_APP_OAUTH_GOOGLE_REDIRECT_URL as string,
      response_type: 'code',
      access_type: 'offline',
      scope: [
        'https://www.googleapis.com/auth/userinfo.profile',
        'https://www.googleapis.com/auth/userinfo.email',
      ].join(' '),
      state,
    })
    loginUrl = `${AUTHORIZE_URI}?${queryStr}`
  }

  return (
    <div className='Home'>
      <a href={loginUrl}>Google Login</a>
    </div>
  )
}

function Home() {
  const [content, setContent] = useState(<div />)

  useEffect(() => {
    axios
      .get(
        `${
          process.env.REACT_APP_SERVER_ORIGIN as string
        }/oauth/google/user-info`,
        {
          withCredentials: true,
        },
      )
      .then((response) => {
        setContent(<div>안녕, {response.data.user_given_name}</div>)
      })
      .catch((error) => {
        console.log(error)
        if (error.response.status === 401) {
          setContent(<Login />)
        }
      })
  }, [])

  return <div className='Home'>{content}</div>
}

export default Home
