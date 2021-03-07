import { UserService } from './../user-service/user.service';
import { Component, OnInit } from '@angular/core';
import { User } from './../models/user';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent implements OnInit {
  displayedColumns: string[] = ['firstName', 'lastName', 'email', 'userName', 'status', 'department', 'tools'];
  users: User[] = [];
  hasError = false;

  constructor(private userService: UserService) { }

  ngOnInit(): void {
    console.log('loading users');
    this.userService.getUsers().subscribe(
      data => { this.processUsers(data); },
      error => { this.handleError(error); }
    );
  }

  private processUsers(users: User[]): void {
    this.users = users;
  }

  private handleError(error: HttpErrorResponse): void {
    if (error.status === 500) {
      this.hasError = true;
    }
  }



}
