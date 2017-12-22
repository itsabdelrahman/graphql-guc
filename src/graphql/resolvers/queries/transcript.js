import { getTranscript } from '../../../datasource';

const transcriptResolver = (obj, args, context) => {
  const { username, password } = context;
  return getTranscript({ username, password });
};

export default transcriptResolver;
