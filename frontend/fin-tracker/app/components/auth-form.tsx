'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { GoogleLogin, CredentialResponse } from "@react-oauth/google"


export default function AuthForm() {
  const [isLogin, setIsLogin] = useState(true)
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const router = useRouter()

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    console.log(e)
  }

  const handleGoogleLogin = (credentials: CredentialResponse) => {
    console.log(credentials)
  }
  const handleGoogleLoginError = () => {
    console.log("error when login")
  }  

  return (
    <Card className="w-[350px]">
      <CardHeader>
        <CardTitle>{isLogin ? 'Iniciar sesión' : 'Registrarse'}</CardTitle>
        <CardDescription>
          {isLogin ? '¿No tienes una cuenta?' : '¿Ya tienes una cuenta?'}
          <Button variant="link" onClick={() => setIsLogin(!isLogin)}>
            {isLogin ? 'Regístrate' : 'Inicia sesión'}
          </Button>
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit}>
          {!isLogin && (
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-4">
                <Label htmlFor="name">Nombre</Label>
                <Input id="name" placeholder="Tu nombre" value={name} onChange={(e) => setName(e.target.value)} required />
              </div>
            </div>
          )}
          <div className="grid w-full items-center gap-4">
            <div className="flex flex-col space-y-4">
              <Label htmlFor="email">Email</Label>
              <Input id="email" type="email" placeholder="tu@email.com" value={email} onChange={(e) => setEmail(e.target.value)} required />
            </div>
          </div>
          <div className="grid w-full items-center gap-4">
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="password">Contraseña</Label>
              <Input id="password" type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
            </div>
          </div>
        </form>
      </CardContent>
      <CardFooter className="flex flex-col space-y-4">
        <Button className="w-full" onClick={handleSubmit}>{isLogin ? 'Iniciar sesión' : 'Registrarse'}</Button>
        <GoogleLogin onSuccess={handleGoogleLogin} onError={handleGoogleLoginError}></GoogleLogin>
      </CardFooter>
    </Card>
  )
}

