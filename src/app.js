import http from 'http';
import express from 'express';
import bodyParser from 'body-parser';
import { makeExecutableSchema } from 'graphql-tools';
import { graphqlExpress, graphiqlExpress } from 'apollo-server-express';
import { graphqlSchema, graphqlResolvers } from './graphql';
import config from './constants/config';

const app = express();
app.server = http.createServer(app);

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

const executableSchema = makeExecutableSchema({
  typeDefs: [graphqlSchema],
  resolvers: graphqlResolvers,
});

app.use(
  '/graphql',
  bodyParser.json(),
  graphqlExpress({
    schema: executableSchema,
  }),
);

app.use(
  '/graphiql',
  graphiqlExpress({
    endpointURL: '/graphql',
  }),
);

app.server.listen(process.env.PORT || config.server.port);
console.log(`🚀  Server listening on port ${app.server.address().port}...`);
