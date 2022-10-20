import bcrypt from 'bcrypt'
import cors from 'cors'
import express, { Request, Response } from 'express'
import { Schema, connect, connection, model } from 'mongoose'

const PORT = 1234
const SALT_ROUNDS = 3
const DB_NAME = 'testDB'
const allowedOrigins = ['http://localhost:3000']
const corsOptions: cors.CorsOptions = {
  origin: allowedOrigins,
}

const app = express()
app.use(express.json())
app.use(cors(corsOptions))

// Connecting MongoDB
connect(`mongodb://localhost:27017/${DB_NAME}`)
const db = connection
db.on('error', function () {
  console.log(`MongoDB Connection Failed! DB name: ${DB_NAME}`)
})
db.once('open', function () {
  console.log(`MongoDB Connected! DB name: ${DB_NAME}`)
})

const isUsernameValid = (username: string) => {
  if (username.length < 6 || username.length > 30) return false
  return true
}

const isPasswordValid = (password: string) => {
  if (password.length < 8 || password.length > 30) return false
  return true
}

// MongoDB interface, schema, model for user
interface IUser {
  username: string
  password: string
}
const userSchema = new Schema<IUser>({
  username: { type: String, required: true },
  password: { type: String, required: true },
})
const User = model<IUser>('User', userSchema)

app.get('/welcome', (req: Request, res: Response) => {
  res.send('welcome!')
})

app.post('/users', async (req: Request, res: Response) => {
  const user = req.body

  // Check a duplicate username
  const existing_user = await User.findOne({ username: user.username })
  if (existing_user !== null)
    return res.status(409).json({ username: user.username })

  if (!isUsernameValid(user.username))
    return res.status(403).json({ message: 'Invalid username' })

  if (!isPasswordValid(user.password))
    return res.status(403).json({ message: 'Invalid password' })

  bcrypt.genSalt(SALT_ROUNDS, (err, salt: string) => {
    if (err) return res.status(500)

    bcrypt.hash(user.password, salt, async (err, hashedPassword: string) => {
      if (err) return res.status(500)

      const new_user = new User({
        username: user.username,
        password: hashedPassword,
      })
      await new_user.save()

      return res.status(200).send(new_user)
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
