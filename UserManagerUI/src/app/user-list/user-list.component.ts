import { UserService } from './../user-service/user.service';
import { Component, OnInit } from '@angular/core';
import { User } from '../user-service/user';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent implements OnInit {
  displayedColumns: string[] = ['firstName', 'lastName', 'email', 'userName', 'status', 'department', 'tools'];
  users: User[];

  constructor(private userService: UserService) { }

  ngOnInit(): void {
    // this.userService.getUsers().subscribe(
    //   data => { this.users = data; },
    // );
    this.userService.getUsers().subscribe(
      data => {
        this.users = data;
      },
      error => { this.handleError(error); }

    );
  }

  private handleError(error: HttpErrorResponse): void {
    if (error.status === 500) {

    }
  }



}
