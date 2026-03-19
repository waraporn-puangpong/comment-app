import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CommentApiModel, CreateCommentRequest } from '../model/comment';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private readonly apiUrl = 'http://localhost:8080/api';

  constructor(private http: HttpClient) {}

  saveComment(comment: CreateCommentRequest): Observable<CommentApiModel> {
    return this.http.post<CommentApiModel>(`${this.apiUrl}/comments/update-comment`, comment);
  }

  getComments(): Observable<CommentApiModel[]> {
    return this.http.get<CommentApiModel[]>(`${this.apiUrl}/comments`);
  }
}
