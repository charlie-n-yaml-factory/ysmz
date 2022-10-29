import cors from 'cors'
import express, { Request, Response } from 'express'
import { connect, connection } from 'mongoose'

import { MONGODB_URI } from './util/secrets'

// Controllers (route handlers)
import * as userController from './controllers/user'

const PORT = 1234
const allowedOrigins = ['http://localhost:3000']
const corsOptions: cors.CorsOptions = {
  origin: allowedOrigins,
}

const app = express()
app.use(express.json())
app.use(cors(corsOptions))

// Connecting MongoDB
connect(MONGODB_URI)
const db = connection
db.on('error', function () {
  console.log('MongoDB Connection Failed!')
})
db.once('open', function () {
  console.log('MongoDB Connected!')
})

app.get('/welcome', (req: Request, res: Response) => {
  res.send('welcome!')
})

app.post('/users', userController.postUser)

app.listen(PORT, () => {
  console.log(`
  #########################################
  🛡️  Server listening on port: ${PORT}  🛡️
  #########################################
`)
})
