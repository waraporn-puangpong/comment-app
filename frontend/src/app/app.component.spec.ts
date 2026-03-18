import { TestBed } from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';

describe('AppComponent', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [RouterTestingModule, FormsModule],
    declarations: [AppComponent]
  }));

  it('should create the app', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });

  it('should seed one comment', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app.comments.length).toBe(1);
  });

  it('should add a trimmed comment and clear draft', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;

    app.draftComment = '  hello  ';
    app.addComment();

    expect(app.comments.length).toBe(2);
    expect(app.comments[1].message).toBe('hello');
    expect(app.comments[1].author).toBe(app.currentUserName);
    expect(app.draftComment).toBe('');
  });
});
