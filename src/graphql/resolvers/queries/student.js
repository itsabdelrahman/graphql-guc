import { GraphQLError } from 'graphql';
import { isAuthorized } from '../../../datasource';

const student = async (obj, args) => {
  const { username, password } = args;
  const areAuthenticCredentials = await isAuthorized({ username, password });

  if (!areAuthenticCredentials) {
    throw new GraphQLError('Unauthorized: Inauthentic credentials');
  }
};

export default student;
