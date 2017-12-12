import { getSchedule } from '../../../datasource';

const scheduleResolver = (obj, args, context) => {
  const { username, password } = context;
  return getSchedule({ username, password });
};

export default scheduleResolver;
