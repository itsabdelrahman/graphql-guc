export default `
schema {
  query: Query
}

type Query {
  student(email: String!, password: String!): Student
}

type Student {
  courses: [Course]
  schedule: [Slot]
}

type Course {
  code: String
  name: String
  absence: Absence
  coursework: [Component]
  midterm: Midterm
  exam: Exam
}

type Absence {
  level: Int
  severity: AbsenceSeverity
}

type Slot {
  type: SlotType
  weekday: SlotWeekday
  number: Int
  venue: String
}

type Component {
  type: String
  grade: Float
  maximumGrade: Float
}

type Midterm {
  grade: Float
}

type Exam {
  venue: String
  seat: String
  startsAt: String
}

enum SlotWeekday {
  SATURDAY
  SUNDAY
  MONDAY
  TUESDAY
  WEDNESDAY
  THURSDAY
}

enum SlotType {
  LECTURE
  TUTORIAL
  LAB
}

enum AbsenceSeverity {
  HIGH
  MEDIUM
  LOW
}
`;
