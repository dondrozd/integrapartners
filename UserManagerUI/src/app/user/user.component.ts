import { UserService } from './../user-service/user.service';
import { User } from './../models/user';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpErrorResponse } from '@angular/common/http';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  newMode = false;
  userId: number = null;
  user: User = {id: null, firstName: '', lastName: '', email: '', userName: '', status: 'A', department: ''};

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private userService: UserService) { }

  ngOnInit(): void {
    if (this.route.snapshot.paramMap.has('id')) {
      this.userId = +this.route.snapshot.paramMap.get('id');
      // load user data
      this.userService.getUser(this.userId).subscribe(
        data => { this.user = data; },
        error => { this.handleError(error); }
      );
    } else {
      this.newMode = true;
    }
  }

  private handleError(error: HttpErrorResponse): void {
    if (error.status === 500) {
      this.router.navigate(['error', '500']);
    }
  }

  onClickSave(): void {
    if (this.newMode) {
      console.log('save new user');
      this.saveNewUser();
    } else {
      console.log('save existing user');
      this.saveExistingUser();
    }
    this.returnToUserList();
  }

  saveNewUser(): void {
    this.userService.addUser(this.user).subscribe(
      data => { this.user = data; },
      error => { this.handleError(error); }
    );
  }

  saveExistingUser(): void {
    this.userService.updateUser(this.user).subscribe(
      data => { this.user = data; },
      error => { this.handleError(error); }
    );
  }


  onClickCancel(): void {
    console.log('cancel');
    this.returnToUserList();
  }


  private returnToUserList(): void {
    this.router.navigate(['users']);
  }

}
