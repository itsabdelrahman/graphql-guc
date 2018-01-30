import { isAuthorized } from '../../../datasource';

const student = async (obj, args, context) => {
  const { username, password } = args;
  const areAuthenticCredentials = await isAuthorized({ username, password });

  // eslint-disable-next-line no-param-reassign
  context.username = username;
  // eslint-disable-next-line no-param-reassign
  context.password = password;

  return { isAuthorized: areAuthenticCredentials };
};

export default student;
