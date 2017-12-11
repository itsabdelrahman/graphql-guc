export default `
type Student {
  isAuthorized: Boolean
}

type Query {
  student(email: String!, password: String!): Student
}

schema {
  query: Query
}
`;
