import {
  student,
  authenticatedStudent,
  courses,
  schedule,
  transcript,
} from './resolvers/queries';
import { login } from './resolvers/mutations';

const resolversMap = {
  Query: {
    student,
    authenticatedStudent,
  },
  Mutation: {
    login,
  },
  Student: {
    courses,
    schedule,
    transcript,
  },
};

export default resolversMap;
