import gql from 'graphql-tag';

export const CREATE_TASK_MUTATION = gql`
  mutation CreateTask($title: String!, $description: String!) {
    insert_tasks(objects: { title: $title, description: $description }) {
      returning {
        id
        title
        description
      }
    }
  }
`;

export const UPDATE_TASK_MUTATION = gql`
  mutation UpdateTask($id: ID!, $title: String!, $description: String!) {
    update_tasks_by_pk(pk_columns: { id: $id }, _set: { title: $title, description: $description }) {
      id
      title
      description
    }
  }
`;

export const DELETE_TASK_MUTATION = gql`
  mutation DeleteTask($id: ID!) {
    delete_tasks_by_pk(id: $id) {
      id
    }
  }
`;
