import React, { useState } from 'react'
import Container from '@mui/material/Container'
import Card from '@mui/material/Card'
import TextField from '@mui/material/TextField'

export default function Home () {
  const [size, setSize] = useState(0)

  const controlSize = (evt: React.ChangeEvent<HTMLInputElement>) => {
    const current: number = Number(evt.target.value)
    setSize(current)
    if (current < 0) {
      setSize(0)
    }
  }

  return (
    <Container sx={{ mt: 10 }}>
      <Card sx={{ mx: 'auto', p: 2, width: 500, height: 500 }}>
        <TextField label="Size" type="number" onChange={ controlSize } value={ size }/>
      </Card>
    </Container>
  )
}
