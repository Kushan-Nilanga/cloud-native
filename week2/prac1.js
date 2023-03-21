const express = require('express');
const app = express();

let users = [
  { id: 1, name: 'John' },
  { id: 2, name: 'Jane' },
  { id: 3, name: 'Bob' }
];

app.use(express.json()); // for parsing application/json <- Add this line

app.get('/users', (req, res) => {
  res.json(users);
});

app.post('/users', (req, res) => {
  const newUser = req.body;
  users.push(newUser);
  res.status(201).json(req.body);
});

app.put('/users/:id', (req, res) => {
  const userId = parseInt(req.params.id);
  const updatedUser = req.body;
  users = users.map(user => user.id === userId ? updatedUser : user);
  res.status(200).json(updatedUser);
});

app.delete('/users/:id', (req, res) => {
  const userId = parseInt(req.params.id);
  users = users.filter(user => user.id !== userId);
  res.status(204).send();
});

app.listen(3000, () => {
  console.log('Server is listening on port 3000');
});


// curl --data "Hello" --request POST localhost:3000/users
// curl -X POST -H "Content-Type:application/json" -d '{ \"key1\": \"value1\" }' http://localhost:3000/users
