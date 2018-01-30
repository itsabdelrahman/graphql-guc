import { getTranscript } from '../../../datasource';

const transcriptResolver = (obj, args, context) => {
  const { isAuthorized } = obj;
  const { username, password } = context;

  if (!isAuthorized) {
    return null;
  }

  return getTranscript({ username, password });
};

export default transcriptResolver;
