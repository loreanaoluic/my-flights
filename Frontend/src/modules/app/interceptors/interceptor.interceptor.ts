import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable()
export class InterceptorInterceptor implements HttpInterceptor {
  
  intercept(
    req: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    const item = localStorage.getItem("userToken");
    if (item) {
      const decodedItem = JSON.parse(item);
      const cloned = req.clone({    
        headers: req.headers.set("Authorization","Bearer " + decodedItem),
      });

      return next.handle(cloned);
    } else {
      return next.handle(req);
    }
  }
}
