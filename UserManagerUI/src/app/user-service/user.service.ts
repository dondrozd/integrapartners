import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from './user';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private usersUrl = 'api/users';

  constructor(private http: HttpClient) { }

  getUsers(): Observable<User[]> {
    return this.http.get<User[]>(this.usersUrl);
  }

  getUser(id: number): Observable<User> {
    return this.http.get<User>(this.usersUrl + '/' + id);
  }

  addUser(user: User): Observable<User> {
    return this.http.post<User>(this.usersUrl + '/' + user.id, user);
  }

  saveUser(user: User): Observable<User> {
    return this.http.put<User>(this.usersUrl + '/' + user.id, user);
  }

  deleteUser(user: User): Observable<User> {
    return this.http.delete<User>(this.usersUrl + '/' + user.id);
  }
}
