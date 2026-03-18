import { Component } from '@angular/core';

type CommentItem = {
  id: number;
  author: string;
  authorInitial: string;
  message: string;
};

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  readonly postAuthor = 'Change can';
  readonly postDateLabel = '16 October 2021 16:00';
  readonly postImageUrl = 'assets/images/images.png';
  readonly currentUserName = 'Blend 285';
  readonly currentUserInitial = 'B';

  draftComment = '';
  comments: CommentItem[] = [
    {
      id: 1,
      author: 'Blend 285',
      authorInitial: 'B',
      message: 'Welcome to my comment section.'
    }
  ];

  private nextCommentId = 2;
  title = 'frontend';

  addComment(): void {
    const message = this.draftComment.trim();
    if (!message) {
      return;
    }

    this.comments = [
      ...this.comments,
      {
        id: this.nextCommentId++,
        author: this.currentUserName,
        authorInitial: this.currentUserInitial,
        message
      }
    ];
    this.draftComment = '';
  }

  trackByCommentId(_: number, comment: CommentItem): number {
    return comment.id;
  }
}
