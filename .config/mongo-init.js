db = db.getSiblingDB("rinha");

db.createCollection("clientes");
db.createCollection("transacoes");

db.clientes.createIndex({ cliente_id: 1 });

db.clientes.insertMany([
  {
    cliente_id: 1,
    limite: 100000,
    saldo: 0,
  },
  {
    cliente_id: 2,
    limite: 80000,
    saldo: 0,
  },
  {
    cliente_id: 3,
    limite: 1000000,
    saldo: 0,
  },
  {
    cliente_id: 4,
    limite: 10000000,
    saldo: 0,
  },
  {
    cliente_id: 5,
    limite: 500000,
    saldo: 0,
  },
]);
