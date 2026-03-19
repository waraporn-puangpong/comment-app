import { Component } from '@angular/core';
import { CreateCommentRequest, CommentApiModel } from './model/comment';
import { ApiService } from './service/api.service';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  ngOnInit() {
    this.getComments();
  }

  title = 'frontend';
  comments: CommentApiModel[] = [];
  constructor(private readonly apiService: ApiService) { }

  readonly postAuthor = 'Change can';
  readonly postImageId = 'post-1';
  readonly postDateLabel = '16 October 2021 16:00';
  readonly postImageUrl = 'assets/images/images.png';
  readonly currentUserName = 'Blend 285';
  readonly currentUserInitial = 'B';

  draftComment = '';
  addComment(): void {
    // เอา message ที่ได้จาก พิมพ์
    const message = this.draftComment

    // mapp data
    const payload: CreateCommentRequest = {
      imageId: this.postImageId,
      username: this.currentUserName,
      content: message,
      createdAt: new Date()
    };

    this.apiService.saveComment(payload).subscribe({
      //คืน response จาก api มาเก็บใน savedComment
      next: (savedComment) => {
        this.draftComment = '';
        this.getComments();
      },
      error: (err: unknown) => {
        console.error('Failed to save comment', err);
      }
    });
  }

  getComments() {
    this.apiService.getComments().subscribe({
      next: (res) => {
        this.comments = res;
      },
      error: (err: unknown) => {
        console.error('Failed to retrieve comments', err);
      }
    });
  }


}