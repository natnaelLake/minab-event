import gql from 'graphql-tag';

export const GET_TASKS_QUERY = gql`
  query GetTasks {
    tasks {
      id
      title
      description
    }
  }
`;

export const GET_TASK_QUERY = gql`
  query GetTask($id: ID!) {
    task(id: $id) {
      id
      title
      description
    }
  }
`;
