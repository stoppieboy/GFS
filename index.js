const express = require('express')
const morgan = require('morgan')

const app = express()

// app.use(morgan('dev'))

app.get('/', (req, res) => {
    res.sendStatus(200)
})

app.listen(3000,() => console.log("server running..."))