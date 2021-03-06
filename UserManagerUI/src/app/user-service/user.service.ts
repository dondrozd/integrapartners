import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from '../models/user';

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
    return this.http.post<User>(this.usersUrl, user);
  }

  updateUser(user: User): Observable<User> {
    return this.http.put<User>(this.usersUrl + '/' + user.id, user);
  }

  deleteUser(id: number): Observable<User> {
    return this.http.delete<User>(this.usersUrl + '/' + id);
  }
}
