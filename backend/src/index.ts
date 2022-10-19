import bcrypt from 'bcrypt'
import cors from 'cors'
import express, { Request, Response } from 'express'

const PORT = 1234
const SALT_ROUNDS = 3
const allowedOrigins = ['http://localhost:3000']
const corsOptions: cors.CorsOptions = {
  origin: allowedOrigins,
}

const app = express()
app.use(express.json())
app.use(cors(corsOptions))

app.get('/welcome', (req: Request, res: Response) => {
  // res.set('Access-Control-Allow-Origin', 'http://localhost:3000')
  res.send('welcome!')
})

app.post('/users', (req: Request, res: Response) => {
  // res.set('Access-Control-Allow-Origin', 'http://localhost:3000')

  const user = req.body

  bcrypt.genSalt(SALT_ROUNDS, (err, salt: string) => {
    if (err) res.status(500)
    bcrypt.hash(user.password, salt, (err, hashedPassword: string) => {
      if (err) res.status(500)
      user.password = hashedPassword
      res.status(200).send(user)
    })
  })
})

app.listen(PORT, () => {
  console.log(`
  #########################################
  🛡️  Server listening on port: ${PORT}  🛡️
  #########################################
`)
})
