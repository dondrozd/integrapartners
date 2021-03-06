import { RouterTestingModule } from '@angular/router/testing';
import { ActivatedRoute } from '@angular/router';
import { UserService } from './../user-service/user.service';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UserComponent } from './user.component';
import { User } from './../models/user';
import { of } from 'rxjs';

describe('UserComponent', () => {

  let component: UserComponent;
  let fixture: ComponentFixture<UserComponent>;
  let userServiceSpy: jasmine.SpyObj<UserService>;

  const user: User = {id: 1, firstName: 'fnA', lastName: 'lnA', email: 'a@dn.com', userName: 'unA', status: 'A', department: 'depA'};

  beforeEach(async () => {
    userServiceSpy = jasmine.createSpyObj(UserService, ['getUser']);
    await TestBed.configureTestingModule({
      imports: [ RouterTestingModule ],
      declarations: [ UserComponent ],
      providers: [
        { provide: UserService, useValue: userServiceSpy },
        { provide: ActivatedRoute, useValue: {snapshot: { paramMap: { get: () => 1, has: () => true } }} }
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UserComponent);
    component = fixture.componentInstance;
  });

  it('should create', () => {
    userServiceSpy.getUser.and.returnValue(of(user));
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });

  // it('should route on cancel', () => {
  //   userServiceSpy.getUser.and.returnValue(of(user));
  //   const location: Location = TestBed.inject(Location);
  //   fixture.detectChanges();
  //   component.onClickCancel();
  //   expect(component).toBeTruthy();
  //   expect(location.pathname).toBe('users');
  // });
});
