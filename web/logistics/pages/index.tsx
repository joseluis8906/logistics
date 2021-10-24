import React, { useState } from 'react'
import Container from '@mui/material/Container'
import Card from '@mui/material/Card'
import TextField from '@mui/material/TextField'
import Box from '@mui/material/Box'
import CardContent from '@mui/material/CardContent'
import { Typography } from '@mui/material'

export default function Home () {
  const guardMinVal = (evt: React.ChangeEvent<HTMLInputElement>, cb: (a: number) => void) => {
    const current: number = Number(evt.target.value)
    if (current >= 0) {
      cb(current)
    }
  }

  const [width, setWidth] = useState(0)
  const controlWidth = (evt: React.ChangeEvent<HTMLInputElement>) => guardMinVal(evt, setWidth)

  const [height, setHeight] = useState(0)
  const controlHeight = (evt: React.ChangeEvent<HTMLInputElement>) => guardMinVal(evt, setHeight)

  const [length, setLength] = useState(0)
  const controlLength = (evt: React.ChangeEvent<HTMLInputElement>) => guardMinVal(evt, setLength)

  const [weight, setWeight] = useState(0)
  const controlWeight = (evt: React.ChangeEvent<HTMLInputElement>) => guardMinVal(evt, setWeight)

  return (
    <Container sx={{ mt: 10 }}>
      <Card sx={{ mx: 'auto', p: 2, width: 500, height: 'auto' }}>
        <CardContent><Typography>Calculate Shipping</Typography></CardContent>
        <CardContent>
          <Box component='form' sx={{ '& > :not(style)': { m: 1, width: '27ch' } }} noValidate autoComplete='off'>
            <TextField label="From" type="text"/>
            <TextField label="To" type="text"/>
            <TextField label="Width" type="number" onChange={ controlWidth } value={ width }/>
            <TextField label="Height" type="number" onChange={ controlHeight } value={ height }/>
            <TextField label="Length" type="number" onChange={ controlLength } value={ length }/>
            <TextField label="Weight" type="number" onChange={ controlWeight } value={ weight }/>
          </Box>
        </CardContent>
      </Card>
    </Container>
  )
}
