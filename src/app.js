import http from 'http';
import express from 'express';
import bodyParser from 'body-parser';
import cors from 'cors';
import { makeExecutableSchema } from 'graphql-tools';
import { graphqlExpress } from 'apollo-server-express';
import graphqlPlayground from 'graphql-playground-middleware-express';
import checkEnv from 'check-env';
import { graphqlSchema, graphqlResolvers } from './graphql';
import config from './constants/config';

checkEnv(['ENCRYPTION_KEY']);

const app = express();
app.server = http.createServer(app);

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(cors());

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
      message: 'Internal Server Error',
      path: error.path,
    }),
    tracing: true,
  }),
);

app.get('/playground', graphqlPlayground({ endpoint: '/graphql' }));

app.get('/', (req, res) => res.redirect('/playground'));

app.server.listen(process.env.PORT || config.server.port);
// eslint-disable-next-line no-console
console.log(`🚀  Server listening on port ${app.server.address().port}...`);
