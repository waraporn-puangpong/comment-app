export interface CreateCommentRequest {
  imageId: string;
  username: string;
  content: string;
  createdAt: Date;
}

export interface CommentApiModel {
  id: string;
  imageId: string;
  username: string;
  content: string;
  createdAt: Date;
}
