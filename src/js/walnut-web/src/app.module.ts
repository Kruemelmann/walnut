import { Module } from '@nestjs/common';
import {ServeStaticModule} from '@nestjs/serve-static';
import {join} from 'path';
import { AppController } from './app.controller';
import { AppService } from './app.service';

@Module({
  imports: [
      ServeStaticModule.forRoot({
//TODO think about a better way to configure a path to the frontend files. Maybe a environment variable so localy i can use direnv and inside a container a can use dockers environment variable.
          rootPath: join(__dirname, '..', 'client'),
      }),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
