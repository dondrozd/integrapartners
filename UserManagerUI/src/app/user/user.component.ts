import { UserService } from './../user-service/user.service';
import { User } from './../models/user';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  newMode = false;
  userId: number = null;
  user: User = {id: 0, firstName: '', lastName: '', email: '', userName: '', status: '', department: ''};

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private userService: UserService) { }

  ngOnInit(): void {
    if (this.route.snapshot.paramMap.has('id')) {
      this.userId = +this.route.snapshot.paramMap.get('id');
      // load user data
      this.userService.getUser(this.userId).subscribe(data => {
        this.user = data;
      });
    } else {
      this.newMode = true;
    }
  }

  onClickSave(): void {
    if (this.newMode) {
      console.log('save new user');
    } else {
      console.log('save existing user');
    }
    this.returnToUserList();
  }


  onClickCancel(): void {
    console.log('cancel');
    this.returnToUserList();
  }


  private returnToUserList(): void {
    this.router.navigate(['users']);
  }

}
