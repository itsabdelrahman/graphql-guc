import { isAuthorized } from '../../../datasource';
import { decryptWithStoredKey } from '../../../utilities';

const authenticatedStudentResolver = async (obj, args, context) => {
  const { token } = args;

  const decryptedCredentials = decryptWithStoredKey(token);
  const [username, password] = decryptedCredentials.split(':');

  const areAuthenticCredentials = await isAuthorized({ username, password });

  // eslint-disable-next-line no-param-reassign
  context.username = username;
  // eslint-disable-next-line no-param-reassign
  context.password = password;

  return { isAuthorized: areAuthenticCredentials };
};

export default authenticatedStudentResolver;
