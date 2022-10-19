import axios from 'axios'
import React from 'react'
import './App.css'

interface User {
  username: string
  password: string
}

function App() {
  const signUp = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()

    const user = {
      username: (event.target as HTMLFormElement).username.value,
      password: (event.target as HTMLFormElement).password.value,
    }

    axios
      .post<User>('http://localhost:1234/users', user, {
        headers: {
          Accept: 'application/json',
        },
      })
      .then((response) => {
        console.log(response)
      })
      .catch((error) => {
        console.log(error)
      })
  }

  return (
    <div className='App'>
      <form onSubmit={signUp}>
        <span>username</span>
        <input type='text' id='username' />
        <br />
        <span>password</span>
        <input type='password' id='password' />
        <br />
        <input type='submit' value='Sign Up' />
      </form>
    </div>
  )
}

export default App
