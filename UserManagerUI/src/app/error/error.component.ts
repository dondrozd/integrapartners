import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-error',
  templateUrl: './error.component.html',
  styleUrls: ['./error.component.css']
})
export class ErrorComponent implements OnInit {
  errorCode = '404';
  errorMessage = 'Page Not Found';

  constructor() { }

  ngOnInit(): void {
  }

}
