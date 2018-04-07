import { isAuthorized } from '../../../datasource';
import { encryptWithStoredKey } from '../../../utilities';

const loginResolver = async (obj, args, context) => {
  const { username, password } = args;
  const areAuthenticCredentials = await isAuthorized({ username, password });

  if (!areAuthenticCredentials) {
    return {
      isAuthorized: false,
      token: null,
    };
  }

  const encryptedCredentials = encryptWithStoredKey(
    [username, password].join(':'),
  );

  return {
    isAuthorized: true,
    token: encryptedCredentials,
  };
};

export default loginResolver;
