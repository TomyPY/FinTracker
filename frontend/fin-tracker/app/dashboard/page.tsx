'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import { Button } from "@/components/ui/button"

export default function Dashboard() {
  const [userName, setUserName] = useState('')
  const router = useRouter()

  useEffect(() => {
    const name = localStorage.getItem('userName')
    if (name) {
      setUserName(name)
    } else {
      router.push('/')
    }
  }, [router])

  const handleLogout = () => {
    localStorage.removeItem('userName')
    router.push('/')
  }

  return (
    <div className="flex min-h-screen flex-col items-center justify-center p-24">
      <h1 className="text-4xl font-bold mb-8">Bienvenido, {userName}!</h1>
      <Button onClick={handleLogout}>Cerrar sesión</Button>
    </div>
  )
}

