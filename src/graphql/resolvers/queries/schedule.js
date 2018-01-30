import { getSchedule } from '../../../datasource';

const scheduleResolver = (obj, args, context) => {
  const { isAuthorized } = obj;
  const { username, password } = context;

  if (!isAuthorized) {
    return null;
  }

  return getSchedule({ username, password });
};

export default scheduleResolver;
