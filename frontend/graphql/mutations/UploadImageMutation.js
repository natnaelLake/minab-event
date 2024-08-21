import gql from 'graphql-tag';

export const UPLOAD_IMAGE = gql`
  mutation UploadImage($file: UploadImageInput!) {
    uploadImage(input: $file) {
      imageUrl
    }
  }
`;