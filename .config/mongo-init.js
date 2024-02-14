db = db.getSiblingDB("rinha");

db.createCollection("customers");
db.createCollection("transactions");

db.customers.createIndex({ customer_id: 1 });
db.transactions.createIndex({ customer_id: 1, created_at: 1 });

db.customers.insertMany([
  {
    customer_id: 1,
    limit: 100000,
    balance: 0,
  },
  {
    customer_id: 2,
    limit: 80000,
    balance: 0,
  },
  {
    customer_id: 3,
    limit: 1000000,
    balance: 0,
  },
  {
    customer_id: 4,
    limit: 10000000,
    balance: 0,
  },
  {
    customer_id: 5,
    limit: 500000,
    balance: 0,
  },
]);
