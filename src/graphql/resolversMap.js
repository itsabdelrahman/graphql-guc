import {
  student,
  courses,
  absence,
  exam,
  schedule,
  transcript,
} from './resolvers/queries';

const resolversMap = {
  Query: {
    student,
  },
  Student: {
    courses,
    schedule,
    transcript,
  },
  Course: {
    absence,
    exam,
  },
};

export default resolversMap;
