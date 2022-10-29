import bcrypt from 'bcrypt'
import { Request, Response } from 'express'
import { User } from '../models/User'

const SALT_ROUNDS = 3

const isUsernameValid = (username: string) => {
  if (username.length < 6 || username.length > 30) return false
  return true
}

const isPasswordValid = (password: string) => {
  if (password.length < 8 || password.length > 30) return false
  return true
}

export const postUser = async (req: Request, res: Response) => {
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
}
