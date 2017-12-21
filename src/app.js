import http from 'http';
import express from 'express';
import bodyParser from 'body-parser';
import { makeExecutableSchema } from 'graphql-tools';
import { graphqlExpress } from 'apollo-server-express';
import graphqlPlayground from 'graphql-playground-middleware-express';
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
    formatError: error => ({
      message: error.message,
      path: error.path,
    }),
  }),
);

app.use('/playground', graphqlPlayground({ endpoint: '/graphql' }));

app.server.listen(process.env.PORT || config.server.port);
// eslint-disable-next-line no-console
console.log(`🚀  Server listening on port ${app.server.address().port}...`);
